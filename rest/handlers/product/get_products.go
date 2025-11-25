package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"net/http"
	"strconv"
)

// var cnt int64
// var mu sync.Mutex

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request){	

 reqQuery := r.URL.Query()

 pageAsStr := reqQuery.Get("page")
 limitAsStr := reqQuery.Get("limit")

 page, _ := strconv.ParseInt(pageAsStr, 10, 32)
 limit, _ := strconv.ParseInt(limitAsStr, 10, 32)
 
 if page <= 0 {
	page = 1
 }

 if limit <= 0 {
	limit = 10
 }
 
 prdCh := make(chan []*domain.Product)
 go func ()  {
   productList, err := h.svc.List(page, limit)
   if err != nil{
	 util.SendError(w, http.StatusInternalServerError, "Internal server error")
	return
 }
   prdCh <- productList
 }()
	
//  var wg sync.WaitGroup
  
 ch := make(chan int64)

//  wg.Add(1)
 go func ()  {
  
	 //defer wg.Done()
   
	 cnt1, err := h.svc.Count()
     if err != nil{
	   util.SendError(w, http.StatusInternalServerError, "Internal server error")
	  return
   }
   
	 ch <- cnt1
 }()

 //wg.Wait()
  // time.Sleep(10 * time.Millisecond)

	//  cnt1, err := h.svc.Count()
  //    if err != nil{
	//    util.SendError(w, http.StatusInternalServerError, "Internal server error")
	//   return
  //  }

	productList := <- prdCh
	totalCnt := <- ch
	
 
 util.SendPage(w, productList, page, limit, totalCnt)


}