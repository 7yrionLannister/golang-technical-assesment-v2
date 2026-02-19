import http from "k6/http";
import { sleep } from "k6";
import { describe } from 'https://jslib.k6.io/k6chaijs/4.3.4.3/index.js';
import { createValidator } from '../helpers/validator.js';
import { SLEEP_DURATION, BASE_URL, TOKEN } from '../helpers/constants.js';

export default function() {
    const validator = createValidator('Consumption API');

    describe("getConsumptionMonthly", () => {
        let param1 = 'online';
        let url = BASE_URL + `/consumption?meter_id=1&start_date=2023-06-01&end_date=2023-06-10&period=daily`;
        let res = http.get(url, {
            headers: {
                'Authorization': `Bearer ${TOKEN}`,
            }
        });
        
        validator(res.status, `getConsumptionMonthly - response status (${res.status})`).to.equal(200);
    });
}