import { Component, OnInit } from '@angular/core';

export class User {
  id: number;
  data: { [key: string]: Object };

  constructor(id: number, data: { [key: string]: Object }) {
    this.id = id;
    this.data = data;
  }
}

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}
