import { Component, OnInit,ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { HttpClient } from '@angular/common/http'
import { MatTableDataSource } from '@angular/material/table';
import { MatDialog } from '@angular/material/dialog';
import { DeleteDialogComponent } from '../delete-dialog/delete-dialog.component';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
@Component({
  selector: 'app-list-user',
  templateUrl: './list-user.component.html',
  styleUrls: ['./list-user.component.css']
})
export class ListUserComponent implements OnInit {
  user:any
  userID:any
  dataColumns:string[]=['name','mail','role','createddate','action'];
  obj = {}
  dataSource:any
  count = 0
  @ViewChild(MatPaginator)
  paginat !: MatPaginator

  constructor(
    private http:HttpClient,
    private dialog: MatDialog,
    private toast: ToastrService,
    private route:Router
    ) { }

  ngOnInit(): void {
    this.userCount()
    this.fetch()
  }
  fetch(){
    this.http.post("http://localhost:3000/listUser",this.obj).subscribe((res:any)=>{
    this.user = res
    this.dataSource =  new MatTableDataSource(this.user)
    this.dataSource.paginator=this.paginat
    })

  }
  userCount(){
    this.http.post("http://localhost:3000/userCount",{}).subscribe((res:any)=>{
    this.count=res
  })
  }

  delete(id:any){
    this.userID=id
    let box=this.dialog.open(DeleteDialogComponent,{width:"500px",disableClose:true,})
    box.afterClosed().subscribe((res:any)=>{
     if (res==true){
      this.http.post("http://localhost:3000/deleteUser",{"id":this.userID}).subscribe((res:any)=>{
        if (res.Verification==true){
            this.toast.success('','Deleted User Successfully')
            this.route.navigate(['/user/listUser'])
        }else{
          this.toast.error('','User not Found')
        }
        this.fetch()
      })
     }
    })
  }
}
