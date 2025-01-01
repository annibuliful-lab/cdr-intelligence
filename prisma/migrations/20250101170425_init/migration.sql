-- CreateEnum
CREATE TYPE "PermissionAbility" AS ENUM ('READ', 'UPDATE', 'DELETE', 'CREATE', 'EXECUTE');

-- CreateTable
CREATE TABLE "accounts" (
    "id" UUID NOT NULL,
    "username" TEXT NOT NULL,
    "password" VARCHAR(256) NOT NULL,
    "createdAt" TIMESTAMPTZ NOT NULL,
    "updatedAt" TIMESTAMPTZ NOT NULL,
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "accounts_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "sub_district" (
    "id" UUID NOT NULL,
    "subDistrictId" TEXT NOT NULL,
    "subDistrictName" TEXT NOT NULL,
    "districtId" TEXT NOT NULL,
    "zipCode" TEXT NOT NULL,
    "provinceProvinceId" TEXT NOT NULL,

    CONSTRAINT "sub_district_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "district" (
    "districtId" TEXT NOT NULL,
    "districtName" TEXT NOT NULL,

    CONSTRAINT "district_pkey" PRIMARY KEY ("districtId")
);

-- CreateTable
CREATE TABLE "province" (
    "provinceId" TEXT NOT NULL,
    "provinceName" TEXT NOT NULL,

    CONSTRAINT "province_pkey" PRIMARY KEY ("provinceId")
);

-- CreateTable
CREATE TABLE "telecommunication_service_provider" (
    "id" UUID NOT NULL,
    "name" VARCHAR(128) NOT NULL,
    "description" TEXT NOT NULL,
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "telecommunication_service_provider_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "telecommunication_service_provider_batching_data_configuration" (
    "id" UUID NOT NULL,
    "name" TEXT NOT NULL,
    "isDefault" BOOLEAN NOT NULL DEFAULT true,
    "serviceProviderId" UUID NOT NULL,
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "telecommunication_service_provider_batching_data_configura_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "BatchingDataFieldConfiguration" (
    "id" UUID NOT NULL,
    "configurationId" UUID NOT NULL,
    "sourceType" TEXT NOT NULL,
    "sourceField" TEXT NOT NULL,
    "targetType" TEXT NOT NULL,
    "targetField" TEXT NOT NULL,

    CONSTRAINT "BatchingDataFieldConfiguration_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "call_detail_records" (
    "id" UUID NOT NULL,
    "serviceProviderId" UUID NOT NULL,
    "sourceMsisdn" TEXT NOT NULL,
    "sourceImsi" TEXT,
    "sourceImei" TEXT,
    "destinationMsisdn" TEXT NOT NULL,
    "destinationImsi" TEXT,
    "destinationImei" TEXT,
    "callDurationInMinute" INTEGER NOT NULL,
    "locationLatitude" DOUBLE PRECISION NOT NULL,
    "locationLongitude" DOUBLE PRECISION NOT NULL,
    "address" TEXT NOT NULL,
    "callStartTime" TIMESTAMP(3),
    "callEndTime" TIMESTAMP(3),
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "call_detail_records_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Permission" (
    "id" UUID NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "Permission_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "projects" (
    "id" UUID NOT NULL,
    "title" TEXT NOT NULL,
    "createdAt" TIMESTAMPTZ NOT NULL,
    "updatedAt" TIMESTAMPTZ NOT NULL,
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "projects_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project_roles" (
    "id" UUID NOT NULL,
    "projectId" UUID NOT NULL,
    "title" VARCHAR(32) NOT NULL,
    "createdAt" TIMESTAMPTZ NOT NULL,
    "updatedAt" TIMESTAMPTZ NOT NULL,
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "project_roles_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project_role_permission" (
    "id" UUID NOT NULL,
    "permissionId" UUID NOT NULL,
    "projectRoleId" UUID NOT NULL,
    "createdAt" TIMESTAMPTZ NOT NULL,
    "updatedAt" TIMESTAMPTZ NOT NULL,
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "project_role_permission_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project_accounts" (
    "id" UUID NOT NULL,
    "projectRoleId" UUID NOT NULL,
    "accountId" UUID NOT NULL,
    "createdAt" TIMESTAMPTZ NOT NULL,
    "updatedAt" TIMESTAMPTZ NOT NULL,
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "project_accounts_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "accounts_username_key" ON "accounts"("username");

-- CreateIndex
CREATE UNIQUE INDEX "Permission_name_key" ON "Permission"("name");

-- CreateIndex
CREATE UNIQUE INDEX "project_role_permission_projectRoleId_permissionId_key" ON "project_role_permission"("projectRoleId", "permissionId");

-- CreateIndex
CREATE UNIQUE INDEX "project_accounts_accountId_projectRoleId_key" ON "project_accounts"("accountId", "projectRoleId");

-- AddForeignKey
ALTER TABLE "sub_district" ADD CONSTRAINT "sub_district_districtId_fkey" FOREIGN KEY ("districtId") REFERENCES "district"("districtId") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "sub_district" ADD CONSTRAINT "sub_district_provinceProvinceId_fkey" FOREIGN KEY ("provinceProvinceId") REFERENCES "province"("provinceId") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "telecommunication_service_provider_batching_data_configuration" ADD CONSTRAINT "telecommunication_service_provider_batching_data_configura_fkey" FOREIGN KEY ("serviceProviderId") REFERENCES "telecommunication_service_provider"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "BatchingDataFieldConfiguration" ADD CONSTRAINT "BatchingDataFieldConfiguration_configurationId_fkey" FOREIGN KEY ("configurationId") REFERENCES "telecommunication_service_provider_batching_data_configuration"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "call_detail_records" ADD CONSTRAINT "call_detail_records_serviceProviderId_fkey" FOREIGN KEY ("serviceProviderId") REFERENCES "telecommunication_service_provider"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project_roles" ADD CONSTRAINT "project_roles_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES "projects"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project_role_permission" ADD CONSTRAINT "project_role_permission_projectRoleId_fkey" FOREIGN KEY ("projectRoleId") REFERENCES "project_roles"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project_role_permission" ADD CONSTRAINT "project_role_permission_permissionId_fkey" FOREIGN KEY ("permissionId") REFERENCES "Permission"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project_accounts" ADD CONSTRAINT "project_accounts_projectRoleId_fkey" FOREIGN KEY ("projectRoleId") REFERENCES "project_roles"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project_accounts" ADD CONSTRAINT "project_accounts_accountId_fkey" FOREIGN KEY ("accountId") REFERENCES "accounts"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
