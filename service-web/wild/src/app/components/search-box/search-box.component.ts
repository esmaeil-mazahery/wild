import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-search-box',
  templateUrl: './search-box.component.html',
  styleUrls: ['./search-box.component.scss'],
})
export class SearchBoxComponent implements OnInit {
  @Output() search = new EventEmitter<string>();
  Term = new FormControl();
  constructor() {}

  ngOnInit(): void {}

  onSearch() {
    this.search.emit(this.Term.value);
  }

  onEnter(event: any) {
    if (event.keyCode == 13 && this.Term.valid) {
      this.Term.markAsDirty();
      this.onSearch();
      return false;
    }
    return true;
  }
}
