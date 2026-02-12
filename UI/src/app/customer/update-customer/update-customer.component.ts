import { Component, OnInit } from '@angular/core';
import { ActivatedRoute,Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { ToastrService } from 'ngx-toastr';
@Component({
  selector: 'app-update-customer',
  templateUrl: './update-customer.component.html',
  styleUrls: ['./update-customer.component.css']
})
export class UpdateCustomerComponent implements OnInit {

  id:any
  customerDetail:any={address:{}}
  constructor(
    private router:ActivatedRoute,
    private route:Router,
    private http:HttpClient,
    private toast:ToastrService
  ) { }

  ngOnInit(): void {
    this.id=this.router.snapshot.paramMap.get('id')
    this.http.post("http://localhost:3000/listOneCustomer",{"id":this.id}).subscribe((res:any)=>{
      console.log(res.Result)
      this.customerDetail=res.Result
    })
  }

  update(){
    this.http.post("http://localhost:3000/updateCustomer",this.customerDetail,{}).subscribe((res:any)=>{
       if (res.Verification==true){
        this.toast.success('','CustomerUpdated Sucessfull')
        this.route.navigate(['/customer/listCustomer'])
       }else{
        this.toast.error('','User Not Found')
       }
    })
  }

}
