go build -o gows . 

sudo cp gows /usr/local/bin/

if [[ "$OSTYPE" == "darwin"* ]]; then
    os_plateform="OSX"
    echo "OSX PLATEFORM detected, gows will run using launchd"
elif [[ "$OSTYPE" == "linux"* ]]; then
    echo "LINUX PLATEFORM detected, gows will run using systemctl"
    os_plateform="LINUX"

else 
    echo "OS not supported by gows"
    exit 1
fi


if [[ "$os_plateform" == "OSX" ]]; then
    gows_daemon=$(launchctl list | grep com.gows.daemon.plist)
    if [ -n "$(launchctl list | grep com.gows.daemon.plist)" ]; then
        echo "gows daemon is already running, restart..."     
        launchctl unload ~/Library/LaunchAgents/com.gows.daemon.plist
    fi

    # Give permission on gows's log dir & application dir
    sudo sh -c "mkdir -p /var/log/gows ; chown -R $USER: /var/log/gows; mkdir -p /var/lib/gows;  chown -R $USER: /var/lib/gows"
    touch /var/log/gows/stdout.log /var/log/gows/stderr.log    
    
    # Add gows daemon to launchctl
    cp daemon/mac/com.gows.daemon.plist  ~/Library/LaunchAgents/
    launchctl load ~/Library/LaunchAgents/com.gows.daemon.plist
fi

if [[ "$os_plateform" == "LINUX" ]]; then
    sudo cp daemon/linux/gows.service /etc/systemd/system/
    if [ -n "$(systemctl list-units --type=service | grep gows)" ]; then
        echo "gows daemon is already running, reload..."
        sudo systemctl disable gows.service
    fi	
    {
        # Make gows unit to start at boot time
	sudo systemctl enable /etc/systemd/system/gows.service  
        sudo systemctl reload-or-restart gows.service
    } || {
    
        echo "Gows is not properly installed" && exit 1
    }

    if [ -n "$(systemctl list-units --type=service --state=active | grep gows)" ]; then
	echo "gows service is now active"
    else 
        echo "gows service fail to start" && exit 1
    fi    		
fi

echo "Installation completed, use 'gows -h' to display gows usage "

