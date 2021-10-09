 const searchButton = document.getElementById("check-availability")
 

  let attention = Prompt()



  searchButton.addEventListener("click",()=>{
    const args = {
  msg:"well hello there",
  icon:"success",
  position: "top-end",
}

    attention.toast(args)
      
  })



  button.addEventListener("click",()=>{


 let html  =`
 <form id="check-dates-form1" action="" method="post"  novalidate class= "needs-validation">
    <div class= "row">
      <div class= "col">
          <div id ="check-dates-form" class= "row check-dates-form">

            <div class= "col">
              <input disabled required class="form-control type= "date" id="start" placeholder="Arrival">
            </div>

            <div class= "col">
              <input  disabled required class="form-control type= "date" id="end" placeholder="Departure">
            </div>


            
          </div>

      </div>
    </div>
 
 </form>
 
 
 `



 const c = {
  msg:html,
  html: html,
  title:"choose your dates"
}
 
  attention.custom(c);




errorButton.addEventListener("click",()=>{

const c = {
  msg:"whoops there was a problem "
}
 
  attention.error(c);

})

  })




const elem2 = document.getElementById("date-picker-test");

const datepicker = new DateRangePicker(elem2, {

  format:"dd/mm/yyyy"
  // ...options
}); 
