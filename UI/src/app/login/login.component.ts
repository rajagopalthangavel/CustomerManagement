import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import {  Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
   mail:string=''
   password:string=''
   hide=true
   body={mail:'',password:'',role:''}
  constructor(
    private http:HttpClient,
    private route:Router,
    private toast:ToastrService) { }

  ngOnInit(): void {
  }
  login(){
    this.body.mail=this.mail
    this.body.password=this.password
    this.body.role="ROLE_ADMIN"
    console.log(this.body)

    this.http.post("http://localhost:3000/login",this.body,{}).subscribe((res:any)=>{
      if (res.Verification==true){
         this.route.navigate(["/customer/listCustomer"])
         this.toast.success('',res.Result)
        }else{
          this.toast.error(res.Error)
        }
    })
  }
}
