# Run Benchmark

You will need to have Python3 installed.

Start up a Python3 virtual environment:
```
python3 -m venv env
source env/bin/activate.fish
```

Install dependencies:
```
pip install -r requirements.txt
```

Seed Nungwi:
```
./seed.py
```

Run Apache Benchmark:
```
ab -n 2000 -c 20 -p body.txt -T application/json http://localhost:8080/nungwi.v1alpha.NungwiService/Check
```
