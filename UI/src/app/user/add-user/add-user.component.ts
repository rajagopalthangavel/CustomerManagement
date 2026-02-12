import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Toast, ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';

@Component({
  selector: 'app-add-user',
  templateUrl: './add-user.component.html',
  styleUrls: ['./add-user.component.css']
})
export class AddUserComponent implements OnInit {

  employeeDetails = { name: '', mail: '', role: '',password:'' };
  hidePass=true
  constructor(private http:HttpClient,private toast:ToastrService,private route:Router) { }

  ngOnInit(): void {
  }
  create(){
    this.http.post("http://localhost:3000/createUser",this.employeeDetails,{}).subscribe((res:any)=>{
      if(res.Verification==true){
        this.toast.success('','Created User Succesfull')
        this.route.navigate(["/user/listUser"])
      }else{
        this.toast.error('','User Creation Error..!!!')
      }
    })
  }
}
