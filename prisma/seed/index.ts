import { seedAccounts } from './account';
import { seedPermissions } from './permission';

async function main() {
  await seedPermissions();
  await seedAccounts();
}

main().then(() => {
  process.exit(0);
});
