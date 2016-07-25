Stuck? Hopefully the hints in here can help you get through your issue,
without giving the answer away.
problems to generate
- make_manifest should require an extra param that's not documented
- make compile script for nginx fail
- make compile script for nginx pull the wrong blob
- make manifest template have a bad port
- make manifest template use a hash instead of an array
- insert a bug in the job start script for nginx 
- insert a bug in the erb for fubar - failure to compile erb
- insert a bug in the erb for nginx config (put logging value in as port value)
- remove package dep for fubar job
- remove package dep for fubar package (golang)
- require static IP for fubar (so nginx can forward), but don't set it up for one (requires multiple jobs)

