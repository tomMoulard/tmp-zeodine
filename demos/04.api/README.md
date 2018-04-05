En utilisant la lib [httprouter](https://github.com/julienschmidt/httprouter), écrire un programme en go qui réponde aux requêtes suivantes :


## GetTime

Retourne la date et l'heure.

```
API : /gettime
PARAMS : néant
RETURN : date et heure format long
```

**Exemple**

```
http://localhost:8080/gettime

2009-11-10 23:00:00 +0000 UTC m=+0.000000001
```

## GetReverse

Renvoi la chaine de caractère à l'envers.

```
API : /getreverse
PARAMS : tout ce qui vient après /getreverse/ est considéré comme une chaine de caractères
RETURN : la chaine de caractères dans le sens inverse
```

**Exemple**
```
http://localhost:8080/getreverse/anystring

gnirtsyna
```

## GetOld

Renvoi l'age en fonction de l'année de naissance.

```
API : /getold
PARAMS : année sur 4 digits, par ex : /getold/1999
RETURN : l'age 
```

**Exemple**
```
http://localhost:8080/getold/1999

19
```

