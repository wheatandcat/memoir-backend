import dayjs from 'dayjs';
import { cd } from 'zx';
import { Timestamp } from 'firebase-admin/firestore';
import { getCoverageTotal } from './coverage.js';
import { sendMetrics } from './client.js';

const add = 0;

const aggregateMetricsAndSend = async () => {
  cd('../../');

  const coverage = await getCoverageTotal();

  const metrics = {
    coverage: coverage.coverageTotal,
    date: Timestamp.fromDate(new Date(dayjs().add(add, 'days').format())),
  };

  await sendMetrics(dayjs().add(add, 'days').format('YYYY-MM-DD'), metrics);
};

const main = async () => {
  await aggregateMetricsAndSend();
};

main();
