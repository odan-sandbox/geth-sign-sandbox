echo abcde > ./out/plain.txt
openssl dgst -sha1 -sign ./out/private.pem ./out/plain.txt > ./out/signature1.dat
openssl dgst -sha1 -sign ./out/private.pem ./out/plain.txt > ./out/signature2.dat

openssl asn1parse -inform DER -in ./out/signature1.dat
openssl asn1parse -inform DER -in ./out/signature2.dat