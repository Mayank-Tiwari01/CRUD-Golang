### Functions that Passed:
1. **Doctor Functions**:
   - `CreateDoctor`
   - `GetDoctorByID`
   - `UpdateDoctor`
   - `DeleteDoctor`
   - `SearchDoctorByName` (Endpoint worked, but test failed)

2. **Patient Functions**:
   - `CreatePatient`
   - `GetPatientByID`

### Functions that Failed:
1. **Doctor Functions**:
   - `TestSearchDoctorByName`
     - **Reason**: The test assertion failed because the response contained more items than expected. The test expected 1 item but found 2 in the search results.