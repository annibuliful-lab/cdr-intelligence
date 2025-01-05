import { v7 } from 'uuid';
import { client } from './client';
import { hash } from 'argon2';

async function seedAccountA() {
  const project = await client.project.create({
    data: {
      id: v7(),
      title: 'Test Project',
      createdBy: 'Seed',
    },
  });

  const role = await client.projectRole.create({
    data: {
      projectId: project.id,
      id: v7(),
      title: 'Test Role',
      createdBy: 'Seed',
    },
  });

  const account = await client.account.create({
    data: {
      id: v7(),
      username: 'testA',
      password: await hash('testA'),
      createdBy: 'Seed',
    },
  });

  const projectAccount = await client.projectAccount.create({
    data: {
      id: v7(),
      projectId: project.id,
      projectRoleId: role.id,
      createdBy: 'Seed',
      accountId: account.id,
    },
  });

  const permissions = await client.permission.findMany({});

  for (const permission of permissions) {
    await client.projectRolePermission.create({
      data: {
        id: v7(),
        permissionId: permission.id,
        projectRoleId: projectAccount.projectRoleId,
        createdBy: 'Seed',
      },
    });
  }
}
export async function seedAccounts() {
  await seedAccountA();
}
