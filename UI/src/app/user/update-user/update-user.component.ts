import { HttpClient } from '@angular/common/http';
import { identifierName } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
@Component({
  selector: 'app-update-user',
  templateUrl: './update-user.component.html',
  styleUrls: ['./update-user.component.css']
})
export class UpdateUserComponent implements OnInit {
  employeeDetails :any= {};
  Id:any
  hidePass=true
  constructor(
    private router:ActivatedRoute,
    private http:HttpClient,
    private toast:ToastrService,
    private route:Router
    ) { }

  ngOnInit(): void {
    this.Id= this.router.snapshot.paramMap.get('id')
    this.http.post("http://localhost:3000/listOneUser",{"id":this.Id}).subscribe((res:any)=>{
    this.employeeDetails=res
    })
  }
  Update(){
   this.employeeDetails.id=this.Id
      this.http.post("http://localhost:3000/updateUser",this.employeeDetails).subscribe((res:any)=>{
        if (res.Verification==true){
          this.toast.success('','Updated User Successfull')
          this.route.navigate(['/user/listUser'])
        }else{
          this.toast.error('','Updated Error.!!!')
        }
      })
  }
}
