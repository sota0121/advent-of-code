# Mojo setup

ref: https://docs.modular.com/mojo/manual/get-started/

## Installation

install magic

```sh
curl -ssL https://magic.modular.com/d9ec4367-ee88-4b69-b783-c3484a32a41e | bash
```

create a mojo project with magic

```sh
magic init hello-world --format mojoproject
```

start a shell in the project virtual environment

```sh
cd hello-world && magic shell
```

after magic shell starting, we can start using Mojo.

```sh
mojo --version
```

## Usage

### REPL

```sh
mojo
```

### Run a Mojo file

```sh
mojo hello.mojo
```

### Build an executable binary

```sh
mojo build hello.mojo

# run
./hello
```

## Examples

```sh
git clone https://github.com/modularml/mojo.git
cd mojo/examples

magic run mojo hello_inteerop.mojo
```
