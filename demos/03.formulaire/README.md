Ecrire un programme en go qui écoute en http sur le port 8080.

Ce programme affiche par défaut un formulaire de contact qui demande le prénom et l'année de naissance + un bouton de confirmation.

```
Votre prénom : [______]
Votre année de naissance : [____]

                            [Confirmer]
```

Lorsque l'utilisateur confirme, le programme go affiche l'âge de la personne comme suit :

```
Bonjour %prenom%.

Vous avez %age% ans.
```

Etendre le programme pour contrôler que l'année de naissance est bien un nombre et renvoi une erreur sinon.

Etendre le programme pour utiliser les templates go pour générer la page de résultat.