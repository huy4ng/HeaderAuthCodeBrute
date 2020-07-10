# HeaderAuthCodeBrute
To brute the "Authorization" key of Http headers for tomcat/nexus and so on
------------------------------------
 Author      |       huy4ng
------------------------------------
 Update-v1.0 |      2020-07-10
-------------------------------------

Usage : 
	Example : ./HeaderAuthCodeBrute -L=url.txt -U=username.txt -P=password.txt -T=5

View more help via ./HeaderAuthCodeBrute -h

Usage of ./HeaderAuthCodeBrute:
  -L string
    	Urls list to brute
  -P string
    	Passwords list to brute
  -T int
    	Set Threads (default 5) (default 5)
  -U string
    	Usernames list to brute
