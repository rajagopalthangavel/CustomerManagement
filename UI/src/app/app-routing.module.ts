import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';

const routes: Routes = [
  {path:'' ,redirectTo:'login',pathMatch:'full'},
  {path:'login',component:LoginComponent},
  {path:'user',loadChildren:()=>import('./user/user.module').then(m => m.UserModule)},
  {path:'customer', loadChildren:()=>import('./customer/customer.module').then(n => n.CustomerModule)}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
