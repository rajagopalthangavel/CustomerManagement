import { Component, OnInit, ViewChild } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { ToastrService} from 'ngx-toastr';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { MatDialog } from '@angular/material/dialog';
import { DeleteDialogComponent } from 'src/app/user/delete-dialog/delete-dialog.component';
import { DatePipe } from '@angular/common';
@Component({
  selector: 'app-list-customer',
  templateUrl: './list-customer.component.html',
  styleUrls: ['./list-customer.component.css']
})



export class ListCustomerComponent implements OnInit {

  @ViewChild(MatPaginator)
  paginate !: MatPaginator

  cus:any
  dataColumns:string[]=['name','mail','phone','createddate','action']
  dataSource:any
  count:any
  date="11/01/2023"
  constructor(
    private http:HttpClient,
    private toast:ToastrService,
    private router:Router,
    private dialog:MatDialog,
    private datepipe:DatePipe
  ) { }

  ngOnInit(): void {
    this.CustomerCount()
    this.CustomerList()
    console.log("Latest Date :",this.date)
    let latest_date =this.datepipe.transform(this.date, 'yyyy-MM-dd');
    console.log("Latest Date :",latest_date)
  }

  CustomerList(){
       this.http.post("http://localhost:3000/listCustomer",{}).subscribe((res:any)=>{
       if (res.Verification==true){
        this.cus=res.Result
        this.dataSource=new MatTableDataSource(this.cus)
        this.dataSource.paginator=this.paginate
       }else{
          this.toast.error('','Credential Error')
       }
      })
  }

  CustomerCount(){
    this.http.post("http://localhost:3000/customerCount",{}).subscribe((res:any)=>{
      if (res.Verification==true){
        this.count=res.Result
       }
     })
  }

  delete(id:any){
      let box=this.dialog.open(DeleteDialogComponent,{width:"500px",disableClose:true})
      box.afterClosed().subscribe((res:any)=>{
        if (res==true){
          this.http.post("http://localhost:3000/deleteCustomer",{"id":id},{}).subscribe((res2:any)=>{
            if (res2.Verification==true){
              this.toast.success('','Deleted Successfull')
              this.router.navigate(['/customer/listCustomer'])
            }else{
              this.toast.error('','User not Found')
            }
            this.CustomerList()
          })
        }
      })
  }

}
