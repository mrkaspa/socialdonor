#Social Donor

It's a platform to share blood donning interests and to find blood donors using facebook notificacions(for this the domain has to use SSL). [Gota de vida](https://gotadevida.co/) is a project that uses this repo. The app uses cartoo db maps to display where are the blood donors located and uses geo referenced queries to find the nearest ones.

![Image of gotadevida](https://github.com/mrkaspa/socialdonor/blob/master/demo.png)

To deploy this project in dokku remember to set these variables

#Config in dokku

dokku config:set app DOMAIN="localhost:9000" MYSQL_DB="user:password@tcp\(mariadb:3306\)/db_name?charset=utf8&parseTime=True" FACEBOOK_APP_ID="xxx" FACEBOOK_APP_SECRET="xxx" CARTOO_DOMAIN="xxx" CARTOO_TOKEN="xxx"
