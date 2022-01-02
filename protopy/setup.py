from setuptools import setup, find_namespace_packages

setup(
    name="seekret.public_protobuf",
    description="The Python package containing Seekret's generated protobuf files",
    version="0.1.0",
    packages=find_namespace_packages(include=['seekret.*']),
    install_requires=[
        'grpcio~=1.33',
        'protobuf~=3.14'
    ]
)
