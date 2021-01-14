## Add/Change new user data
1. First fork this repo.
2. Clone forked repo
```
$ git clone https://github.com/${your-username}/MemBi
$ cd MemBi
```
3. See format of toml https://github.com/hahwul/MemBi/blob/main/data/README.md
4. Add or Change user data in `./data/*.toml`
+ e.g (hahwul.toml)
```toml
name = "hahuwl"
site = "https://www.hahwul.com"
hackerone = "hahwul"
bugcrowd = "hahwul"
twitter = "hahwul"
github = "hahwul"
instagram = "hahwul___"
youtube = "https://youtube.com/c/하훌"
stackoverflow = "11547708/hahwul"
```
5. git push 
```
$ git add ./data/*.toml ; git commit -m "change toml data" ; git push
```
6. Make and Send Pull-Request for master repo

## Note
`./distribute` lets you reflect it in README. However, this operation is performed automatically when pushing through Github-action.

* Q. I want to see the readme. What if I want to do this? 🤔
* A. 🚀
```
$ go build distribute.go
$ ./distribute
```
