package handler

type Handler struct {
	DataBase *DataBaseHandler
	Cache    *CacheHandler
	Queue    *QueueHandler
}
