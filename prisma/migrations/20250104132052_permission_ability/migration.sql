/*
  Warnings:

  - You are about to drop the column `action` on the `Permission` table. All the data in the column will be lost.
  - Added the required column `ability` to the `Permission` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Permission" DROP COLUMN "action",
ADD COLUMN     "ability" "PermissionAbility" NOT NULL;
