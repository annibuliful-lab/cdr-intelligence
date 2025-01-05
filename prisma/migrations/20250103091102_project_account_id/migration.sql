/*
  Warnings:

  - Added the required column `projectId` to the `project_accounts` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "project_accounts" ADD COLUMN     "projectId" UUID NOT NULL;

-- AddForeignKey
ALTER TABLE "project_accounts" ADD CONSTRAINT "project_accounts_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES "projects"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
