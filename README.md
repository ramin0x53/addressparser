# addressparser

Build:
```
go build main
```

Example:

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
```