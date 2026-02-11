import tenantsTest from './cases/v1/tenant.js';
import deviceTest from './cases/v1/inventory/device.js';
import networkTest from './cases/v1/inventory/network.js';
import interfaceTest from './cases/v1/inventory/interface.js';
import componentTest from './cases/v1/inventory/component.js';
import entityTest from './cases/v1/inventory/entity.js';
import configurationTest from './cases/v1/inventory/configuration.js';
import pollerTest from './cases/v1/pollers/poller.js';
import pollerHistoryTest from './cases/v1/pollers/poller-history.js';
import statisticsTest from './cases/v1/statistics.js';
import alertTest from './cases/v1/alert/alert.js';
import alertHistoryTest from './cases/v1/alert/alert-history.js';
import credentialTest from './cases/v1/credential.js';
import usageTest from './cases/v1/usage.js';
import asmTest from './cases/v1/asm.js';
import deviceV2Test from './cases/v2/inventory/device.js';
import interfaceV2Test from './cases/v2/inventory/interface.js';

export default function() {
    tenantsTest();
    deviceTest();
    networkTest();
    interfaceTest();
    componentTest();
    entityTest();
    configurationTest();
    pollerTest();
    pollerHistoryTest();
    statisticsTest();
    alertTest();
    alertHistoryTest();
    credentialTest();
    usageTest();
    asmTest();
    deviceV2Test();
    interfaceV2Test();
}
