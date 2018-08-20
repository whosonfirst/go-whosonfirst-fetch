# go-whosonfirst-fetch

## Important

This will get merged / reconciled with the [go-whosonfirst-clone](https://github.com/whosonfirst/go-whosonfirst-clone) and [go-whosonfirst-bundles](https://github.com/whosonfirst/go-whosonfirst-bundles) repos... eventually.

## Tools

### wof-fetch-ids

For example:

```
./bin/wof-fetch-ids -target data -reader 'type=github repo=whosonfirst-data' -reader 'type=github repo=whosonfirst-data-postalcode-us' 102527513
2018/08/20 16:11:04 fetch 102527513 from https://raw.githubusercontent.com/whosonfirst-data/whosonfirst-data/master/data/102/527/513/102527513.geojson and write to data/102/527/513/102527513.geojson
2018/08/20 16:11:04 do fetch for data/102/527/513/102527513.geojson : true
2018/08/20 16:11:05 fetch 85922583 from https://raw.githubusercontent.com/whosonfirst-data/whosonfirst-data/master/data/859/225/83/85922583.geojson and write to data/859/225/83/85922583.geojson
2018/08/20 16:11:05 do fetch for data/859/225/83/85922583.geojson : true
2018/08/20 16:11:05 fetch 102087579 from https://raw.githubusercontent.com/whosonfirst-data/whosonfirst-data/master/data/102/087/579/102087579.geojson and write to data/102/087/579/102087579.geojson
2018/08/20 16:11:05 do fetch for data/102/087/579/102087579.geojson : true
2018/08/20 16:11:05 fetch 102191575 from https://raw.githubusercontent.com/whosonfirst-data/whosonfirst-data/master/data/102/191/575/102191575.geojson and write to data/102/191/575/102191575.geojson
2018/08/20 16:11:05 do fetch for data/102/191/575/102191575.geojson : true
2018/08/20 16:11:05 fetch 102085387 from https://raw.githubusercontent.com/whosonfirst-data/whosonfirst-data/master/data/102/085/387/102085387.geojson and write to data/102/085/387/102085387.geojson
2018/08/20 16:11:05 do fetch for data/102/085/387/102085387.geojson : true
2018/08/20 16:11:05 fetch 85633793 from https://raw.githubusercontent.com/whosonfirst-data/whosonfirst-data/master/data/856/337/93/85633793.geojson and write to data/856/337/93/85633793.geojson
2018/08/20 16:11:05 do fetch for data/856/337/93/85633793.geojson : true
2018/08/20 16:11:05 fetch 85688637 from https://raw.githubusercontent.com/whosonfirst-data/whosonfirst-data/master/data/856/886/37/85688637.geojson and write to data/856/886/37/85688637.geojson
2018/08/20 16:11:05 do fetch for data/856/886/37/85688637.geojson : true
2018/08/20 16:11:05 fetch 554784711 from https://raw.githubusercontent.com/whosonfirst-data/whosonfirst-data-postalcode-us/master/data/554/784/711/554784711.geojson and write to data/554/784/711/554784711.geojson
2018/08/20 16:11:05 do fetch for data/554/784/711/554784711.geojson : true
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-readwrite
