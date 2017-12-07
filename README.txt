Install protobuf
mkdir tmp
cd tmp
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install

find ./ -name .git* -d | xargs rm -fr
