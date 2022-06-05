# addressparser

## Prerequisites

Before using the Go bindings, you must install the libpostal C library. Make sure you have the following prerequisites:

**On Ubuntu/Debian**
```
sudo apt-get install curl autoconf automake libtool pkg-config
```

**On CentOS/RHEL**
```
sudo yum install curl autoconf automake libtool pkgconfig
```

**On Mac OSX**
```
sudo brew install curl autoconf automake libtool pkg-config
```

**Installing libpostal**

```
git clone https://github.com/openvenues/libpostal
cd libpostal
./bootstrap.sh
./configure --datadir=[...some dir with a few GB of space...]
make
sudo make install

# On Linux it's probably a good idea to run
sudo ldconfig
```


## Build:
```
git clone https://github.com/ramin0x53/addressparser.git
cd ./addressparser
make build
```
## Docker:
```
docker build -t addressparser-img .
docker run -d -p 80:80 addressparser-img
```

## Example:

```
Request:

{
    "address":"The Book Club 100-106 Leonard St Shoreditch London EC2A 4RH, United Kingdom"
}


Response:

{
  "city": "london",
  "country": "united kingdom",
  "house": "the book club",
  "house_number": "100-106",
  "postcode": "ec2a 4rh",
  "road": "leonard st",
  "suburb": "shoreditch"
}


Default port: 80
stats: /stats
```
