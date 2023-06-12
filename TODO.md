# Log structure
- Abstract from file?  Might be useful if we want to make this for general use, but easy enough to change later.
- Most of the categories are basically the same and can use the same record type for simplicity only exception is vinnies and other might not have pages, but thats what null is for anyways
- Thought keeping track of completed might be neat to have especially as an optional field, "how many vinnies have i finished?" also potential integration with future systems
- Would a volume variable be useful? probably not but throwing it out there 
- We could require chars/pages or time to be non-null depending on case, but cant see how allowing people to create empty records would break anything

Possible schema
```	
input:
	Category [subgroups] mandatory_arg (optional_arg)
	Reading [VN, Book, Manga, Nuki, Other] name=null (chars=null) (pages=null)  (time=null) (date=now) (completed=0)
	epeen:char/350, page, manga/nuki pages/5,some time factor that will be abused so probably not
```

Example:
```
		/log VN ぴことちこショタアイドルのオシゴト 100000
		Stores the following entry in database
		VN; ぴことちこショタアイドルのオシゴト;100000;null;null;now;0
  
          general reading stats would aggregate VN, book, manga, nuki, other
          custom format for each type would save a few bytes, but is that a huge issue?
          simple case reading would be an alternate if you think it will be
```
 ```       
			Listening [TV, Anime, other]	(name=null) (time=null) (episodes) (date=now) (completed=0)	
      Epeen:Time/episodes
	Storage
		some kind of tsv/csv?
    object solution?
 ```  
 
# Data sources

VNDB API, replace name with VNDB id.

Bookmeter API replace name with bookmeter ID

TVDB has an API, replace name with id.

exhentai/nhentai do they have apis?

calculate times from episodes length and amount for watching time.

# Leaderboards

Achievements

Weekly/Monthly Goals 

Basic point system, ideally would be waited based on effort as opposed to being heavily time biased.

Stats graphs.

# VNDB ideas

get tag information from completed VNS and use it for a recommendation system.

integrate VN of the month and give points once the VN is completed on VNDB,
