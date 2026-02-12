import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
@Component({
  selector: 'app-add-customer',
  templateUrl: './add-customer.component.html',
  styleUrls: ['./add-customer.component.css']
})
export class AddCustomerComponent  {
  customerDetail:any={name:'',mail:'',address:{area:'',city:'',pincode:'',phone:''}}

  constructor(
    private http:HttpClient,
    private toast:ToastrService,
    private router:Router
    ) { }


  CreateCus(){
    console.log(this.customerDetail)
    this.http.post("http://localhost:3000/addCustomer",this.customerDetail,{}).subscribe((res:any)=>{
      if (res.Verification==true){
          this.toast.success('','Customer Created Successful')
          this.router.navigate(['/customer/listCustomer'])
      }else{
        this.toast.error('','Customer Creation Error')
      }
    })
  }
}
