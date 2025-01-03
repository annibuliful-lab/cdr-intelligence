-- AlterTable
ALTER TABLE "projects" ADD COLUMN     "description" TEXT;

-- CreateTable
CREATE TABLE "SessionToken" (
    "token" TEXT NOT NULL,
    "accountId" UUID NOT NULL,
    "revoke" BOOLEAN NOT NULL DEFAULT false,
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "SessionToken_pkey" PRIMARY KEY ("token")
);

-- AddForeignKey
ALTER TABLE "SessionToken" ADD CONSTRAINT "SessionToken_accountId_fkey" FOREIGN KEY ("accountId") REFERENCES "accounts"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
