#!/usr/bin/python3 

from pathlib import Path

def main():
    working_directory = Path(__file__).parent.absolute()
    relative_env_location = '../../.env'
    env_location = working_directory.joinpath(relative_env_location).absolute()
    replace_in_file(
        input_file=working_directory.joinpath('realm_template.json').absolute(),
        output_file=working_directory.joinpath('realm.json').absolute(),
        env_params=read_env_file(env_location))

def read_env_file(path):
    with open(path, 'r') as f:
        lines = f.readlines()

    res = {}
    for line in lines:
        parts = line.split('=', maxsplit=1)
        if not len(parts) == 2:
            continue
        key, value = parts
        key = key.strip()
        value = value.strip()

        if value.startswith('"') and value.endswith('"'):
            value = value[1:-1]
        res[key] = value
    return res

def replace_env_params(content, env_params):
    for key, value in env_params.items():
        content = content.replace('$'+key, value)
    return content

def replace_in_file(input_file, output_file, env_params):
    with open(input_file, 'r') as f:
        content = f.read()
    content = replace_env_params(content, env_params)
    with open(output_file, 'w') as f:
        f.write(content)

main()