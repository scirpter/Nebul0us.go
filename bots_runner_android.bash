ARCH=$(dpkg --print-architecture)

FILE_NAME="bots_android_"

if [ "$ARCH" = "aarch64" ]; then
    FILE_NAME+="arm64-v8a"
elif [ "$ARCH" = "arm" ]; then
    FILE_NAME+="armeabi-v7a"
elif [ "$ARCH" = "x86_64" ]; then
    FILE_NAME+="x86_64"
elif [ "$ARCH" = "i686" ]; then
    FILE_NAME+="x86"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

if [ -f "bots.zip" ]; then
    unzip bots.zip
    rm -rf bots.zip
    if [ -d "$HOME/bots" ]; then
        # if theres tokens.txt in bots folder, copy it over to storage/downloads
        if [ -f "$HOME/bots/tokens.txt" ]; then
            cp -f ~/bots/tokens.txt ~/storage/downloads
        fi
	# if theres auth.key in bots folder, copy it over to storage/downloads
        if [ -f "$HOME/bots/auth.key" ]; then
            cp -f ~/bots/auth.key ~/storage/downloads
        fi
        rm -rf ~/bots
	sleep 1
    fi
    mv -f bots ~
fi

if [ -d "bots" ]; then
    mv -f bots ~
fi

# move tokens.txt back into bots folder
if [ -f "tokens.txt" ]; then
    rm -rf ~/bots/tokens.txt
    mv -f tokens.txt ~/bots
fi

# move auth.key back into bots folder
if [ -f "auth.key" ]; then
    rm -rf ~/bots/auth.key
    mv -f auth.key ~/bots
fi

cd ~
cd bots
chmod 777 $FILE_NAME
./$FILE_NAME
