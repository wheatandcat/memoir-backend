import { initializeApp, cert } from 'firebase-admin/app';
import { getFirestore, Timestamp } from 'firebase-admin/firestore';
import { createRequire } from 'module';
const require = createRequire(import.meta.url);
const serviceAccount = require('./serviceAccount.json');

type MetricItem = {
  coverage: number;
  date: Timestamp;
};

initializeApp({
  credential: cert(serviceAccount),
});

const db = getFirestore();

export const sendMetrics = async (key: string, value: MetricItem) => {
  const docRef = db.collection('backend-metrics').doc(key);
  await docRef.set(value);
};
