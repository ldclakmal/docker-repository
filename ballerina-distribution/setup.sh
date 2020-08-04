mkdir repo
mkdir home

echo "# Set environment variables"
export packageUser=$1
export packagePAT=$2

echo "# Clone GitHub repositories"
cd repo
git clone --depth 1 https://github.com/ballerina-platform/ballerina-lang.git
git clone --depth 1 https://github.com/ballerina-platform/ballerina-distribution.git

echo "# Build ballerina distribution"
cd ballerina-distribution
./gradlew build -x test -x testExamples -x testStdlibs

echo "# Unzip ballerina distribution"
cd ballerina/build/distributions
unzip ballerina-linux-*.zip
mv ballerina-linux-*/distributions/ballerina-* /ballerina/home

echo "# Ballerina version"
cd /
ln -s /ballerina/home/ballerina-*/bin/ballerina /usr/bin/ballerina
ballerina -v
