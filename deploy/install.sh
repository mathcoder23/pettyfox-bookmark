echo install redis
git clone -b docker-redis-6.0.6 https://github.com/mathcoder23/scripts.git docker-redis-6.0.6
cd docker-redis-6.0.6
chmod +x ./install.sh
./install.sh
cd ..
