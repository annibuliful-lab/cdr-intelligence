/*
  Warnings:

  - You are about to drop the `BatchingDataFieldConfiguration` table. If the table is not empty, all the data it contains will be lost.
  - A unique constraint covering the columns `[subject,ability,name]` on the table `Permission` will be added. If there are existing duplicate values, this will fail.

*/
-- DropForeignKey
ALTER TABLE "BatchingDataFieldConfiguration" DROP CONSTRAINT "BatchingDataFieldConfiguration_configurationId_fkey";

-- DropIndex
DROP INDEX "Permission_name_key";

-- DropTable
DROP TABLE "BatchingDataFieldConfiguration";

-- CreateTable
CREATE TABLE "batching_data_field_configuration" (
    "id" UUID NOT NULL,
    "configurationId" UUID NOT NULL,
    "sourceType" TEXT NOT NULL,
    "sourceField" TEXT NOT NULL,
    "targetType" TEXT NOT NULL,
    "targetField" TEXT NOT NULL,

    CONSTRAINT "batching_data_field_configuration_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Permission_subject_ability_name_key" ON "Permission"("subject", "ability", "name");

-- AddForeignKey
ALTER TABLE "batching_data_field_configuration" ADD CONSTRAINT "batching_data_field_configuration_configurationId_fkey" FOREIGN KEY ("configurationId") REFERENCES "telecommunication_service_provider_batching_data_configuration"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
