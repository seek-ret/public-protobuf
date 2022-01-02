#! /usr/bin/python3

import sys
import requests


def del_tag(token, package_id):
    resp = requests.post("https://api.github.com/graphql", headers={
        "Authorization": f"bearer {token}",
        "Accept": "application/vnd.github.package-deletes-preview+json"
    }, json={"query": f'mutation {{ deletePackageVersion(input:{{packageVersionId:"{package_id}"}}) {{ success }}}}'})

    try:
        success = resp.json()["data"]["deletePackageVersion"]["success"]
    except:
        assert False, f"Failed deleting {resp.content}"
    assert success, f"Failed deleting {resp.content}"


def main():
    expected_tag = sys.argv[1]
    token = sys.argv[2]

    resp = requests.post("https://api.github.com/graphql", headers={
        "Authorization": f"bearer {token}",
        "Accept": "application/vnd.github.packages-preview+json"
    }, json={"query": "{organization(login: \"seek-ret\") { packages(packageType: NPM, first: 100) { edges { node { name id versions(first: 100) { nodes {id version}}}}}}}"})

    edges = resp.json()["data"]["organization"]["packages"]["edges"]
    protobuf_versions = [i["node"] for i in edges if i["node"]['name'] == "public-protobuf"][0]['versions']["nodes"]
    for entry in protobuf_versions:
        if entry["version"] == expected_tag:
            del_tag(token, entry["id"])


if __name__ == "__main__":
    sys.exit(main())
