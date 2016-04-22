# Running

```sh
$ docker run -it --rm -v "$PWD":/usr/src/app -w /usr/src/app java java -cp checkstyle-6.17-all.jar com.puppycrawl.tools.checkstyle.Main -c /google_checks.xml Main.java
```