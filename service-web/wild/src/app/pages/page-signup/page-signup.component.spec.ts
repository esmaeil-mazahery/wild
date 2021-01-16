import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PageSignupComponent } from './page-signup.component';

describe('PageSignupComponent', () => {
  let component: PageSignupComponent;
  let fixture: ComponentFixture<PageSignupComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PageSignupComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PageSignupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
