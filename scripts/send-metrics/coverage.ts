import { $ } from 'zx';

export const getCoverageTotal = async () => {
  let r = {
    coverageTotal: 0
  };


  try {
    const result =
      await $`go tool cover -func=coverage.out`;

    const totalLine = result.stdout.split('\n').find(line => line.includes('total:')) || ""
    const coverageTotal = totalLine.split('\t').at(-1)?.replace('%', '') || 0
    
    r.coverageTotal = Number(coverageTotal)

  } catch (err) {
    console.error(err);
  }

  return r;
};
