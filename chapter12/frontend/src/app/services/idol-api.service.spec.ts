import { TestBed } from '@angular/core/testing';

import { IdolApiService } from './idol-api.service';

describe('IdolApiService', () => {
  let service: IdolApiService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(IdolApiService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
