# Welcome to Buffalo!

Thank you for choosing Buffalo for your web development needs.


## Database Setup

It looks like you chose to set up your application using a postgres database! Fantastic!

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start postgres for you.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started postgres, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a

### Migrate Your Databases

The database was created succefuly, but no table yet. Buffalo can do it for you as well: 

	$ buffalo db migrate

## Install node modules
Since Buffalo uses webpack for asset pipeline management, you need to install the required node modules before starting your  application:
```
$ npm install
```

If you don't have NodeJS installed on your machine you need it first.
### Installing NodeJS on macOS
- Install Homebrew if you don't have it (but you really *should*)
```
$ /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```
- Then install NodeJS
```
$ brew install node
```

### Installing NodeJS on Ubuntu/Debian
```
$ curl -sL https://deb.nodesource.com/setup_8.x | sudo -E bash -
$ sudo apt-get install -y nodejs
```


## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)
