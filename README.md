# logstream #

[![Python Versions](https://img.shields.io/pypi/pyversions/pypistats.svg?logo=python&logoColor=FFE873)](https://pypi.org/project/pypistats/)
[![Package Version](https://img.shields.io/pypi/v/logstream.svg)](https://pypi.python.org/pypi/logstream/)

logstream is basically "tail -f" on operating system syslog and log files.  
logstream is a command line tool and python generator.  

### Run from source ###

```bash
./src/logstream/logstream.py
```

### Install via pip ###

```bash
pip install logstream
```

```text
Usage: logstream

    --help
    --version
    --format [type]

    type darwin:
        [default|compact|json|ndjson|syslog]

    type linux:
        [short|short-full|short-unix|verbose|export]
        [json|json-pretty|json-sse|json-seq]

    --tail file

```

---

### logstream in a python shell ###

```python3
>>> import logstream
```

logstream is a generator

```python3
>>> logstream.stream()
<generator object stream at 0x7f806026d0b0>
```

stream is a yeild of each line

```python3
>>> for line in logstream.stream():
...     print(line)
```

tail a file

```
>>> logstream.tail('/tmp/file')
<generator object tail at 0x7f806034b2e0>
>>> for line in logstream.tail('/tmp/file'):
...     print(line)
```

