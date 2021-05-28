go build -o gows . && cp gows /usr/local/bin/

sudo mkdir -p /var/log/gows
sudo chown -R $USER: /var/log/gows 
touch /var/log/gows/stdout.log /var/log/gows/stderr.log 

if [[ "$OSTYPE" == "darwin"* ]]; then
    os_plateform="OSX"
elif [[ "$OSTYPE" == "linux"* ]]; then
    os_plateform="LINUX"
fi

if [[ "$os_plateform" == "OSX" ]]; then
    gows_daemon=$(launchctl list | grep com.gows.daemon.plist)
    if [ -n "$(launchctl list | grep com.gows.daemon.plist)" ]; then
        echo "gows daemon is already running, restart..."     
        launchctl unload ~/Library/LaunchAgents/com.gows.daemon.plist
    fi    
    cp daemon/mac/com.gows.daemon.plist  ~/Library/LaunchAgents/
    launchctl load ~/Library/LaunchAgents/com.gows.daemon.plist
fi

if [[ "$os_plateform" == "LINUX" ]]; then
    cp daemon/linux/gows.service /etc/systemd/system/
    systemctl start gows
fi

