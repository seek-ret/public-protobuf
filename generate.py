import argparse
import glob
import os
import posixpath
import re
import subprocess
import sys

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
PROTO_DIR = os.path.join(SCRIPT_DIR, 'proto')

TYPESCRIPT_OUTPUT_DIR = os.path.join('.', 'protots')
TYPESCRIPT_PROTOC_PLUGIN = os.path.join('.bin',
                                        'protoc-gen-ts.cmd' if sys.platform == 'win32' else 'protoc-gen-ts')
NODEJS_GRPC_PLUGIN = os.path.join('.bin',
                                  'grpc_tools_node_protoc_plugin.cmd' if sys.platform == 'win32' else 'grpc_tools_node_protoc_plugin')
PROTO_FILES = glob.glob(os.path.join(
    PROTO_DIR, '**', '*.proto'), recursive=True)


def generate_python():
    subprocess.check_call(
        [sys.executable, os.path.join(SCRIPT_DIR, 'protopy', 'make.py')])


def generate_typescript(node_modules):
    subprocess.check_call(['protoc',
                           f'--plugin=protoc-gen-grpc={os.path.join(node_modules, NODEJS_GRPC_PLUGIN)}',
                           f'--plugin=protoc-gen-ts={os.path.join(node_modules, TYPESCRIPT_PROTOC_PLUGIN)}',
                           f'--js_out=import_style=commonjs,binary:{TYPESCRIPT_OUTPUT_DIR}',
                           f'--ts_out=service=grpc-node,mode=grpc-js:{TYPESCRIPT_OUTPUT_DIR}',
                           f'--grpc_out=grpc_js:{TYPESCRIPT_OUTPUT_DIR}',
                           f'-I={PROTO_DIR}',
                           *PROTO_FILES],
                          cwd=SCRIPT_DIR)

    # Delete empty GRPC generated files.
    for grpc_generated_file in glob.iglob(os.path.join(TYPESCRIPT_OUTPUT_DIR, '**', '*_grpc_pb.js'), recursive=True):
        with open(grpc_generated_file) as f:
            if f.read().strip() == '// GENERATED CODE -- NO SERVICES IN PROTO':
                f.close()
                os.remove(grpc_generated_file)


_GO_PACKAGE_STATEMENT_REGEX = re.compile(r'option\s+go_package\s*=\s*"([^"]*)"')
_SEEKRET_PROTOBUF_ROOT_GO_MODULE = 'github.com/seek-ret/public-protobuf'


def go_package_name(proto_file: str) -> str:
    with open(proto_file) as f:
        for line in f:
            groups = _GO_PACKAGE_STATEMENT_REGEX.match(line)
            if groups:
                return groups.group(1)

    raise ValueError(f'{proto_file!r} does not contain go package declaration')


def generate_go():
    for file in PROTO_FILES:
        go_package = go_package_name(file)
        if not go_package.startswith(_SEEKRET_PROTOBUF_ROOT_GO_MODULE):
            raise ValueError(
                f'{go_package!r} is not under seekret protobuf root {_SEEKRET_PROTOBUF_ROOT_GO_MODULE!r}')
        package_subpath = go_package[len(
            _SEEKRET_PROTOBUF_ROOT_GO_MODULE):].lstrip('/')

        subprocess.check_call(['protoc',
                               f'-I={PROTO_DIR}',
                               f'--go_out={posixpath.join(".", package_subpath)}', f'--go_opt=module={go_package}',
                               f'--go-grpc_out={posixpath.join(".", package_subpath)}', f'--go-grpc_opt=module={go_package}',
                               file],
                              cwd=SCRIPT_DIR)


def main():
    parser = argparse.ArgumentParser(os.path.basename(sys.argv[0]),
                                     description="Generates code for all the protobuf files, for all of the supported "
                                                 "languages (currently, Go, TypeScript, and Python).")
    parser.add_argument(
        "--node-modules", default=f"{SCRIPT_DIR}/node_modules", dest="node_modules")
    parser.add_argument("-a", "--all", action="store_true")
    parser.add_argument("--py", action="store_true", help="compiles python")
    parser.add_argument("--ts", action="store_true",
                        help="compiles typescript")
    parser.add_argument("--go", action="store_true", help="compiles go")

    args = parser.parse_args()

    if args.all or args.go:
        generate_go()
    if args.all or args.ts:
        generate_typescript(args.node_modules)
    if args.all or args.py:
        generate_python()


if __name__ == '__main__':
    main()
