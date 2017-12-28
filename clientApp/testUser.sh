echo ""
echo "sending the signup, response:"
curl -X POST http://127.0.0.1:3130/signup -d '{"email": "user1@e.com", "password": "user1"}'

echo ""
echo "sending the login, response:"
curl -X POST http://127.0.0.1:3130/login -d '{"email": "user1@e.com", "password": "user1"}'


echo ""
echo "send pubK and m to blind sign"
echo "json to send to the serverIDsigner:"
echo '{"pubKstring": {"e": "65537", "n": "139093"}, "m": "hola"}'
echo "serverIDsigner response:"
BLINDSIGNED=$(curl -X POST http://127.0.0.1:3130/blindsign -d '{"pubKstring": {"e": "65537", "n": "139093"}, "m": "hola"}')
echo "$BLINDSIGNED"

echo ""
echo "send blindsigned to the serverIDsigner to verify"
curl -X POST http://127.0.0.1:3130/verifysign -d '{"m": "hola", "mSigned": "131898 40373 107552 34687"}'
