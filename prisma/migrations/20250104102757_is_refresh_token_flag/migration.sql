-- AlterTable
ALTER TABLE "SessionToken" ADD COLUMN     "isRefreshToken" BOOLEAN NOT NULL DEFAULT false;
