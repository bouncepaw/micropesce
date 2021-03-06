micropesce
##########

**micropesce** is an implementation of some of *pesce* concepts; if you know who i am, you probably know what *pesce* is.

**micropesce** might be an experiment to see how good or bad *pesce* as a concept is while still being somewhat useful. one could describe **micropesce** as a monolithic web-interfaced *vk*-fetching *pesce*.

here's what i want **micropesce** to have:

+ an interface where you can specify what *vk* feeds to fetch and, perhaps, how;
+ an interface for incoming posts;
+ an interface for postponed posts;
+ an interface for saved posts: tags, collections.
   
all four sections can be done as tabs in an application. the application shall be web-based. yeah, web is not where the power lies, but i don't want to spend years polishing my native application development skills for an experimental software.

here's the stack i want to use:

+ *golang*, because it is a bearable programming language;
+ stdlib templates, instead of *qtpl* i use in *mycorrhiza*;
+ some *vk* lib maybe;
+ *sqlite3*. it is time to learn db.

if the **micropesce** project goes well and ends up being enjoyable to use, it may be extended to support silos other than *vk*: *gemini*, *rss* andor *atom* andor *json feed*, public *telegram* channels, *gopher*.

the original idea was to fetch *gemini* feeds but then I thought that fetching *vk* is more useful due to vk being a branch of babylon.