import { PermissionAbility } from '@prisma/client';
import { client } from './client';
import { v7 } from 'uuid';

const abilities = Object.values(PermissionAbility);

const subjects = ['project'];

export async function seedPermissions() {
  for (const subject of subjects) {
    for (const ability of abilities) {
      await client.permission.upsert({
        update: {},
        create: {
          id: v7(),
          subject,
          ability,
          name: subject,
        },
        where: {
          subject_ability_name: {
            subject,
            ability,
            name: subject,
          },
        },
      });
    }
  }
}
