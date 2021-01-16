import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PageSigninComponent } from './page-signin.component';

describe('PageSigninComponent', () => {
  let component: PageSigninComponent;
  let fixture: ComponentFixture<PageSigninComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PageSigninComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PageSigninComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
