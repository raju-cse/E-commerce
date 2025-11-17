package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
	"sync"
)

var cnt int64

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
 
 productList, err := h.svc.List(page, limit)
 if err != nil{
	util.SendError(w, http.StatusInternalServerError, "Internal server error")
	return
 }	

 var wg sync.WaitGroup

 wg.Add(1)
 go func ()  {
  
	 defer wg.Done()

	 cnt1, err := h.svc.Count()
     if err != nil{
	   util.SendError(w, http.StatusInternalServerError, "Internal server error")
	  return
   }
   cnt = cnt1
 }()

 wg.Wait()
  // time.Sleep(10 * time.Millisecond)

	//  cnt1, err := h.svc.Count()
  //    if err != nil{
	//    util.SendError(w, http.StatusInternalServerError, "Internal server error")
	//   return
  //  }
 
 util.SendPage(w, productList, page, limit, cnt)


}