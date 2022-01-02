"""
Make script for creating the Seekret proto package wheel file.

Usage: python make.py

The resulting wheel file will be created under the "out" directory in the directory this script is in.
"""
import glob
import os
import shutil
import subprocess
import sys

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
REPO_ROOT = os.path.dirname(SCRIPT_DIR)
VENV_NAME = 'tmp-build'
VENV_DIR = os.path.join(SCRIPT_DIR, VENV_NAME)
VENV_PYTHON = (os.path.join(VENV_DIR, 'Scripts', 'python') if sys.platform == 'win32'
               else os.path.join(VENV_DIR, 'bin', 'python'))
NAMESPACE_PACKAGE_ROOT = os.path.join(SCRIPT_DIR, 'seekret')


def validate_package_structure(module_path):
    non_namespace_part = os.path.relpath(
        module_path, os.path.join(SCRIPT_DIR, 'seekret'))

    compiled = NAMESPACE_PACKAGE_ROOT
    for component in non_namespace_part.split(os.path.sep):
        compiled = os.path.join(compiled, component)
        create_empty_file(os.path.join(compiled, '__init__.py'))


def create_empty_file(path):
    open(os.path.join(SCRIPT_DIR, path), 'w').close()


def main():
    subprocess.check_call(
        [sys.executable, '-m', 'virtualenv', VENV_NAME], cwd=SCRIPT_DIR)

    try:
        subprocess.check_call([VENV_PYTHON, '-m', 'pip', 'install', '-r',
                               os.path.join(SCRIPT_DIR, 'build-requirements.txt'), "--trusted-host=pypi.org"])

        shutil.rmtree(os.path.join(SCRIPT_DIR, 'seekret', 'proto'))

        input_files = [os.path.relpath(f, REPO_ROOT) for f in glob.iglob(
            os.path.join(REPO_ROOT, 'proto', '**', '*.proto'), recursive=True)]
        args = [VENV_PYTHON, '-m', 'grpc_tools.protoc', *input_files,
                '-I=proto',
                f'--python_out={SCRIPT_DIR}',
                f'--grpc_python_out={SCRIPT_DIR}']

        subprocess.check_call(args, cwd=REPO_ROOT)

        for in_file in input_files:
            out_file = os.path.join(SCRIPT_DIR, os.path.relpath(
                in_file, os.path.join(REPO_ROOT, 'proto')))
            validate_package_structure(os.path.dirname(out_file))

        # Delete empty GRPC generated files.
        for grpc_generated_file in glob.iglob(os.path.join(SCRIPT_DIR, '**', '*_pb2_grpc.py'), recursive=True):
            with open(grpc_generated_file) as f:
                if f.read().strip().endswith('import grpc'):
                    f.close()
                    os.remove(grpc_generated_file)
    finally:
        shutil.rmtree(VENV_DIR, ignore_errors=True)


if __name__ == '__main__':
    main()
